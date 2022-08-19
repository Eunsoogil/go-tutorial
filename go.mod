module tutorial

go 1.19

replace internal/greeting => ./internal/greeting

require (
	github.com/lib/pq v1.10.6
	internal/greeting v0.0.0-00010101000000-000000000000
)
