package handler
//
//import (
//	"fmt"
//	"github.com/gomodule/redigo/redis"
//	"time"
//
//
//)
//
//
//
//
//// HandlerLuckyChars 处理kafka消息,将uinString作为hashtag，防止数据倾斜
//func HandlerLuckyChars(ctx *srfs.Context, kafkaData lightUGCRoom.LightUGCRoomRecordKafka) error {
//
//
//	script := "local m=tonumber(redis.call('GET',KEYS[1]) or 0) " +
//		"local n=tonumber(ARGV[1]) " +
//		"local oldTime=redis.call('HGET',KEYS[2],ARGV[2])" +
//		"if(n>m)" +
//		"then redis.call('HINCRBY',KEYS[2],ARGV[2],ARGV[3])" +
//		"  redis.call('INCRBY',KEYS[1],ARGV[3])" +
//		" redis.call('EXPIRE',KEYS[1],ARGV[4])   " +
//		" redis.call('HSET',KEYS[2],ARGV[5],ARGV[6])" +
//		"return {redis.call('HGET',KEYS[2],ARGV[2]),oldTime} " +
//		"else return {-1,-1}  " +
//		"end"
//	timeArray, err := redis.Int64s(redisCli.Do(ctx, "EVAL", script, 2, key1, key2,
//		LuckyStoreCache.DateDuration, kafkaData.LuckyID, kafkaData.TogetherDuration, 60*60*24, comm.LatestUpdateTime,
//		time.Now().Unix()))
//	if err != nil {
//		ctx.Error("脚本执行失败，err=%+v\t key1==%s\tkey2==%s", err, key1, key2)
//		return err
//	}
//	ctx.Debug("操作成功\t key1==%s\tkey2==%s\t timeArray=%d", key1, key2, timeArray)
//	if len(timeArray) != 2 {
//		return nil
//	}
//	if timeArray[0] == -1 && timeArray[1] == -1 {
//		ctx.Debug("今天增加的时间已经满了")
//		return nil
//	}
//	if length, ok := IsNeedToReport(ctx, kafkaData.LuckyID, timeArray[0], timeArray[1]); ok {
//		Report(ctx, kafkaData, timeArray[0], 2, length)
//	}
//	return nil
//}
//
////
//
//
//
//
//// HandlerExpiredLuckyChars 一天没听歌时减少幸运字符的时间
//func HandlerExpiredLuckyChars(ctx *srfs.Context, kafkaData lightUGCRoom.LightUGCRoomRecordKafka) error {
//	redisCli := redis.New(comm.REDISCONFIGURATION)
//	uinString := comm.GetUinString(kafkaData.UinList)
//	luckyCharsKey := comm.GenerateKey(comm.LuckyChars, uinString)
//	redisData, err := redis.Int64s(redisCli.Do(ctx, "HMGET", luckyCharsKey, comm.LatestUpdateTime, kafkaData.LuckyID))
//	if err != nil || len(redisData) != 2 {
//		ctx.Error("从redis中获取数据失败：err==%+v\t luckyCharsKey=%s\t latestUpdateTime=%s\t luckyId=%d ",
//			err, luckyCharsKey, comm.LatestUpdateTime, kafkaData.LuckyID)
//		return fmt.Errorf("从redis中获取数据失败：err==%+v\t luckyCharsKey=%s\t latestUpdateTime=%s\t luckyId=%d ",
//			err, luckyCharsKey, comm.LatestUpdateTime, kafkaData.LuckyID)
//	}
//	// 如果没过期
//	if !comm.IsExpired(redisData[0]) {
//		ctx.Debug("没过期")
//		return nil
//	}
//	remainTime := redisData[1]
//	// 判断剩余时间是否够扣，不能为负数
//	if remainTime < int64(LuckyStoreCache.ReduceDuration) {
//		remainTime = 0
//	} else {
//		remainTime = remainTime - int64(LuckyStoreCache.ReduceDuration)
//	}
//	// LUA脚本
//	script := "local m=tonumber(redis.call('HGET',KEYS[1],ARGV[1]) ) " +
//		"local n=tonumber(ARGV[2]) " +
//		"if(n==m)" +
//		"then redis.call('HMSET',KEYS[1],ARGV[1],ARGV[3])" +
//		"  end "
//
//	_, evalErr := redisCli.Do(ctx, "EVAL", script, 1, luckyCharsKey, kafkaData.LuckyID, redisData[1], remainTime)
//	if evalErr != nil {
//		ctx.Error("往redis中插入数据错误，err=%+v\t luckyCharsKey=%s\t latestUpdateTime=%s\t luckyId=%d\t remainTime=%d",
//			evalErr, luckyCharsKey, comm.LatestUpdateTime, kafkaData.LuckyID, remainTime)
//	}
//	ctx.Debug(" luckyCharsKey=%s\t latestUpdateTime=%d\t luckyId=%d\t remainTime=%d",
//		luckyCharsKey, redisData[0], kafkaData.LuckyID, remainTime)
//	if length, ok := IsNeedToReport(ctx, kafkaData.LuckyID, remainTime, redisData[1]); ok {
//		Report(ctx, kafkaData, remainTime, 4, length)
//	}
//	return evalErr
//
//}
//
