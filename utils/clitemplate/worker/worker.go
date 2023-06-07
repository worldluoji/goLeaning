package worker

type Worker interface {
	Do(dest string) bool
}
