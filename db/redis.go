package db

import (
    "encoding/json"
    "fmt"
    "github.com/tiny-ci/core/types"
    "github.com/go-redis/redis/v7"
)

type RedisDatabase struct {
    client *redis.Client
}

func New(addr string) (RedisDatabase, error) {
    rdb := RedisDatabase{
        client: redis.NewClient(&redis.Options{
            Addr: addr,
        }),
    }

    return rdb, rdb.client.Ping().Err()
}

func (r RedisDatabase) topLevelJobRepr(pipeId string, i int) string {
    return fmt.Sprintf("pipe:%s:job:%d", pipeId, i)
}

func (r RedisDatabase) NewJobInfoRepr(pipeId string, key string, i int) string {
    return fmt.Sprintf("%s:%s", r.topLevelJobRepr(pipeId, i), key)
}

func (r RedisDatabase) NewGitInfoRepr(pipeId string, key string) string {
    return fmt.Sprintf("pipe:%s:git:%s", pipeId, key)
}

func (r RedisDatabase) Populate(ntf *types.ApiNotification, jobs *[]types.Job) error {
    pi := ntf.PipelineId
    item := make(map[string]interface{})

    item[r.NewGitInfoRepr(pi, "repo_url")]    = ntf.Info.URL
    item[r.NewGitInfoRepr(pi, "ref_name")]    = ntf.Info.RefName
    item[r.NewGitInfoRepr(pi, "commit_hash")] = ntf.Info.CommitHash
    item[r.NewGitInfoRepr(pi, "is_tag")]      = fmt.Sprint(ntf.Info.IsTag)

    var jobList []string
    for i, job := range *jobs {
        steps, err := json.Marshal(job.Steps)
        if err != nil { return err }

        item[r.NewJobInfoRepr(pi, "name", i)]  = job.Name
        item[r.NewJobInfoRepr(pi, "image", i)] = job.Image
        item[r.NewJobInfoRepr(pi, "steps", i)] = steps

        jobList = append(jobList, r.topLevelJobRepr(pi, i))
    }

    res := r.client.MSet(item)
    if res.Err() != nil { return res.Err() }

    intRes := r.client.RPush("job:queue", jobList)
    return intRes.Err()
}
