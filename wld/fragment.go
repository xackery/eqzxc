package wld

// Fragment is what every wld object type adheres to
type Fragment interface {
	// FragmentType identifies the fragment type
	FragmentType() string
}
