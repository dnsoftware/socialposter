package poster

type Platform interface {
	SimplePost(content string) error    // простой пост
	PostWithImage(content string) error // пост с картинкой
}

type Poster struct {
	vk *Platform
}

func NewPoster(vk Platform) *Poster {

	poster := &Poster{
		vk: &vk,
	}

	return poster
}
