package db

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/tiny-ci/core/types"
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

func (r RedisDatabase) newJobInfoRepr(pipeId string, key string, i int) string {
	return fmt.Sprintf("%s:%s", r.topLevelJobRepr(pipeId, i), key)
}

func (r RedisDatabase) newGitInfoRepr(pipeId string, key string) string {
	return fmt.Sprintf("pipe:%s:git:%s", pipeId, key)
}

func (r RedisDatabase) Populate(ntf *types.ApiNotification, jobs *[]types.Job) error {
	pi := ntf.PipelineId
	item := make(map[string]interface{})

	item[r.newGitInfoRepr(pi, "repo_url")] = ntf.Info.URL
	item[r.newGitInfoRepr(pi, "ref_name")] = ntf.Info.RefName
	item[r.newGitInfoRepr(pi, "commit_hash")] = ntf.Info.CommitHash
	item[r.newGitInfoRepr(pi, "is_tag")] = fmt.Sprint(ntf.Info.IsTag)

	var jobList []string
	for i, job := range *jobs {
		steps, err := json.Marshal(job.Steps)
		if err != nil {
			return err
		}

		item[r.newJobInfoRepr(pi, "name", i)] = job.Name
		item[r.newJobInfoRepr(pi, "engine", i)] = job.Engine
		item[r.newJobInfoRepr(pi, "image", i)] = job.Image
		item[r.newJobInfoRepr(pi, "steps", i)] = steps

		jobList = append(jobList, r.topLevelJobRepr(pi, i))
	}

	res := r.client.MSet(item)
	if res.Err() != nil {
		return res.Err()
	}

	intRes := r.client.RPush("job:queue", jobList)
	return intRes.Err()
}
