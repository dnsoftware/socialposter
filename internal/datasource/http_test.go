package datasource

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dnsoftware/socialposter/internal/config"
	"github.com/dnsoftware/socialposter/internal/poster"
)

func Test_GetData(t *testing.T) {
	cfg, err := config.NewConfig()
	require.NoError(t, err)

	ds := NewDataSource(cfg)
	posts, err := ds.GetData()
	fmt.Println(posts)
	require.NoError(t, err)
}

func Test_VkPostSimple(t *testing.T) {
	cfg, err := config.NewConfig()
	require.NoError(t, err)

	vk := poster.NewPlatformVk(cfg.VkToken)

	post := poster.NewPoster(vk)
	_ = post

}
