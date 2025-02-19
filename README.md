# titok-plus

微服务式titok后台系统

主要技术栈golang,mysql,redis,gorm,mq,go-zero


**项目背景**

- 对于单体结构的升级补充，
- 以及对于架构上面的进行优化
- 以及处理数据方面的进行优化，


**全部接口**

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ec412087c9e34796b10dcf896c24d9da~tplv-k3u1fbpfcp-watermark.image?)


**项目的架构设计**

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/b1dcd2b6d5d24b438edfb2dd70604c89~tplv-k3u1fbpfcp-watermark.image?)

> **视频接口的优化**

#### 视频流feed的信息优化

- **对于视频的feed流，进行了处理优化，由于接口的处理数据是从数据库取出至多30个的video，同时还包含用视频作者的信息。**
- **常用的方法是通过for循环查看视频作者的用户Id,然后在rpc到用户数据库查询，为了加快查询的速度，这里使用了mapreduce的编程模型，进行了批处理处理。**
- **运用了go-zero里面自带的mapreduce工具,首先对mapper任务进行输入，输入了30个视频信息，同时创建一个map，key表示视频id,value 表示输入的顺序，等待所有mapper输入完成后，多协程进行mapper工作流的处理，查询用户的信息，最后进行单协程的进行reduce整理数据顺序，相比传统的处理方法优化了处理速度，**

####  视频点赞的优化
- **由于视频点赞为高频操作，如果直接走数据库会对数据库造成很大的压力，这里通过redis进行缓存处理，所有的点赞操作会不经过数据库，会直接在缓存处理，然后后续通过mq的异步处理，相关数据，以及后续进行mq定时任务进行缓存和数据之间的同步。**

- **保证数据库数据的一致性，在缓存点赞的同时，由于视频点赞对应着用户喜爱列表，单纯一个redis hash结构存储用户是否点赞不够，还得处理喜爱列表数据库和redis缓存不一致的问题，为应对用户喜爱列表和缓存不一致的问题，还需要创建一个缓存用户点赞和取消点赞的列表(set),单用户查询喜爱列表时候，先从缓存中查找点赞列表和取消点赞列表，通过数据库和缓存的结合，去重和更新进行查找，确保数据库和缓存的一致性。**

- **视频点赞数量的问题，用户查看到的时候是直接从数据库查看的视频点赞数量，结合点赞数量对于数据的时效和缓存的数据一致性的要求不高，且点赞的高频性质，对于视频点赞数量的处理主要是在放回feed流的时候直接从数据库中查找，然后后续mq定时任务同步到数据库中。redis在点赞的时候会创建一个视频变换量，点赞加一，取消点赞减一，**

**-   数据库和redis的消息一致性的问题，通过消息队列进行异步的定时任务，保持和数据库的数据一致性的问题，**

> **用户接口的优化**

- **对于关注这个高频操作，同样可以和视频流点赞功能进行类比，通过缓存来缓解数据库压力**

-   **由于朋友的认证是两两相互关注为朋友，在点击关注操作的时候，使用了mq消息队列进行异步的关系判断，进行建立朋友和失去朋友的关系操作，**

- **对于关注操作，采用和视频类似的方法，先从缓存查看是否有数据存在，如果存在，改变缓存的状态，如果没有从数据库中查找，然后在将添加缓存数据，过后进行定时任务的同步缓存和数据库的数据，根据操作是关注还是取消关注发送异步请求建立好友列表的请求**
  
> 注意关注表和粉丝表共用一个数据库。

- **关注列表：和缓存配合和数据库一起查找返回，**
- **粉丝列表：和缓存配合和数据库一起返回，**
- **朋友列表：直接返回数据库。**
  
   
> **消息接口的优化**

-   **消息接口由于前端系统选用了轮询的方法来查看聊天记录，这样会使得数据库的压力备增，因为会有大部分的查询记录是无效的，所以使用了，配合redis缓存的方式来进行优化处理，首先redis创建hash结构，存下两个用户的聊天id,如果两个用户的id 分别为1，2，则存在缓存的id为（1-2）和（2-1），最新当用户发送信息时候，会将他们的聊天id进行检查，如果存在，则把属于他的hash 的value改为最新消息的时间戳**。
-   **这样当用户进行轮询查看历史记录的时候，会比较缓存里面的两个聊天id的时间戳是否相同，如果相同，则无需查看数据库，如果不相同，则会从缓存读取当前用户未更新的时间戳，从数据库根据这个时间戳读取最新的聊天记录，然后在把缓存的时间戳更新最新为同步状态，为了防止缓存的堆积，为每个聊天id设置时间，如果超过时间没有更新，则会把缓存删除。**

