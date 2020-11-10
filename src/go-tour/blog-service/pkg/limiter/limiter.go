package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	// 获取对应的限流器的键值对名称
	Key(c *gin.Context) string
	// 获取令桶牌
	GetBucket(key string) (*ratelimit.Bucket, bool)
	// 新增多个令桶牌
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
