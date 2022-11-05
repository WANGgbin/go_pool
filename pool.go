package go_pool

import (
	"time"
)

const (
	defaultMaxNum     int = 100000
	defaultMaxIdleDur     = 1 * time.Minute
)

type PoolCfg struct {
	maxNum           int
	workerMaxIdleDur time.Duration // 回收空闲时长大于等于该值的 worker
}

func NewPoolCfg() *PoolCfg {
	return &PoolCfg{
		maxNum:           defaultMaxNum,
		workerMaxIdleDur: defaultMaxIdleDur,
	}
}

func (cfg *PoolCfg) SetMaxNum(poolMaxNum int) {
	cfg.maxNum = poolMaxNum
}

func (cfg *PoolCfg) SetWorkerMaxIdleDur(poolWorkerMaxIdleDur time.Duration) {
	cfg.workerMaxIdleDur = poolWorkerMaxIdleDur
}

type Pool struct {
	cfg *PoolCfg
}

func NewPoolDefault() *Pool {
	return &Pool{
		cfg: NewPoolCfg(),
	}
}

func NewPoolWithCfg(cfg *PoolCfg) {
	return &Pool{
		cfg: cfg,
	}
}

type fn func(args interface{}) interface{}

type Task struct {
	f    fn
	args interface{}
}

func (p *Pool) SubmitTask(task *Task) error {

}
