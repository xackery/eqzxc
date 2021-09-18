package fragment

// Fragment is what every fragment object type adheres to
type Fragment interface {
	// FragmentType identifies the fragment type
	FragmentType() string
}
