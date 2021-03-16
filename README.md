# ES-AutoClean

自动清理 ELK 架构中 elasticseach 的索引

[![dockeri.co](https://dockeri.co/image/mmhk/es-autoclean)](https://hub.docker.com/r/mmhk/es-autoclean)

## 配置文件

```json
{
    "es-endpoint": "${ES_ENDPOINT}", 
    "index_prefix": "${INDEX_PREFIX}",
    "keep_day": ${KEEP_DAY},
    "check_cron": "${CRON_SPEC}"
}
```

参数说明


|Name|ENV|description|default value|
|-|-|-|-|
|es-endpoint|ES_ENDPOINT|elasticseach访问入口|http://127.0.0.1:9200/|
|index_prefix|INDEX_PREFIX|elasticseach索引前缀| |
|keep_day|KEEP_DAY|保留X天内的索引|15|
|check_cron|CRON_SPEC|执行检查的 [cron 表达式](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/tutorial-lesson-06.html) `秒 分 时 日 月 星期`| 0 9 * * * ? |


### How to run

```bash
docker-compose up
```