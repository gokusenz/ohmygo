package ohmygo

func bmi(cm, kg float64, ch chan float64) {
	ch <- kg / ((cm / 100) * (cm / 100))
}
