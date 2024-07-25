package poster

import (
	"github.com/SevereCloud/vksdk/v3/api"
)

type VkPlatform struct {
	vk *api.VK
}

func NewPlatformVk(vkToken string) *VkPlatform {
	vk := api.NewVK(vkToken)

	platform := &VkPlatform{
		vk: vk,
	}

	return platform
}

func (v *VkPlatform) SimplePost(content string) error {

	return nil
}

func (v *VkPlatform) PostWithImage(content string) error {

	return nil
}
