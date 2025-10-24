package scene

import "context"

// Generator 场景生成器
type Generator interface {
	GenerateScene(ctx context.Context, description string, characters []string) (*Scene, error)
}

// Scene 场景定义
type Scene struct {
	ID          string
	Description string
	Location    string
	TimeOfDay   string
	Weather     string
	Characters  []string
	Objects     []string
	ImagePath   string
}

// DefaultGenerator 默认场景生成器
type DefaultGenerator struct{}

// NewGenerator 创建场景生成器
func NewGenerator() Generator {
	return &DefaultGenerator{}
}

// GenerateScene 生成场景
func (g *DefaultGenerator) GenerateScene(ctx context.Context, description string, characters []string) (*Scene, error) {
	// TODO: 实现场景生成逻辑
	scene := &Scene{
		Description: description,
		Characters:  characters,
	}
	return scene, nil
}
