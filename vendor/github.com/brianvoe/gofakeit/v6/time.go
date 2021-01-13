package gofakeit

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Date will generate a random time.Time struct
func Date() time.Time { return date(globalFaker.Rand) }

// Date will generate a random time.Time struct
func (f *Faker) Date() time.Time { return date(f.Rand) }

func date(r *rand.Rand) time.Time {
	return time.Date(year(r), time.Month(number(r, 1, 12)), day(r), hour(r), minute(r), second(r), nanoSecond(r), time.UTC)
}

// DateRange will generate a random time.Time struct between a start and end date
func DateRange(start, end time.Time) time.Time { return dateRange(globalFaker.Rand, start, end) }

// DateRange will generate a random time.Time struct between a start and end date
func (f *Faker) DateRange(start, end time.Time) time.Time { return dateRange(f.Rand, start, end) }

func dateRange(r *rand.Rand, start time.Time, end time.Time) time.Time {
	return time.Unix(0, int64(number(r, int(start.UnixNano()), int(end.UnixNano())))).UTC()
}

// NanoSecond will generate a random nano second
func NanoSecond() int { return nanoSecond(globalFaker.Rand) }

// NanoSecond will generate a random nano second
func (f *Faker) NanoSecond() int { return nanoSecond(f.Rand) }

func nanoSecond(r *rand.Rand) int { return number(r, 0, 999999999) }

// Second will generate a random second
func Second() int { return second(globalFaker.Rand) }

// Second will generate a random second
func (f *Faker) Second() int { return second(f.Rand) }

func second(r *rand.Rand) int { return number(r, 0, 59) }

// Minute will generate a random minute
func Minute() int { return minute(globalFaker.Rand) }

// Minute will generate a random minute
func (f *Faker) Minute() int { return minute(f.Rand) }

func minute(r *rand.Rand) int { return number(r, 0, 59) }

// Hour will generate a random hour - in military time
func Hour() int { return hour(globalFaker.Rand) }

// Hour will generate a random hour - in military time
func (f *Faker) Hour() int { return hour(f.Rand) }

func hour(r *rand.Rand) int { return number(r, 0, 23) }

// Day will generate a random day between 1 - 31
func Day() int { return day(globalFaker.Rand) }

// Day will generate a random day between 1 - 31
func (f *Faker) Day() int { return day(f.Rand) }

func day(r *rand.Rand) int { return number(r, 1, 31) }

// WeekDay will generate a random weekday string (Monday-Sunday)
func WeekDay() string { return weekDay(globalFaker.Rand) }

// WeekDay will generate a random weekday string (Monday-Sunday)
func (f *Faker) WeekDay() string { return weekDay(f.Rand) }

func weekDay(r *rand.Rand) string { return time.Weekday(number(r, 0, 6)).String() }

// Month will generate a random month string
func Month() string { return month(globalFaker.Rand) }

// Month will generate a random month string
func (f *Faker) Month() string { return month(f.Rand) }

func month(r *rand.Rand) string { return time.Month(number(r, 1, 12)).String() }

// Year will generate a random year between 1900 - current year
func Year() int { return year(globalFaker.Rand) }

// Year will generate a random year between 1900 - current year
func (f *Faker) Year() int { return year(f.Rand) }

func year(r *rand.Rand) int { return number(r, 1900, time.Now().Year()) }

// TimeZone will select a random timezone string
func TimeZone() string { return timeZone(globalFaker.Rand) }

// TimeZone will select a random timezone string
func (f *Faker) TimeZone() string { return timeZone(f.Rand) }

func timeZone(r *rand.Rand) string { return getRandValue(r, []string{"timezone", "text"}) }

// TimeZoneFull will select a random full timezone string
func TimeZoneFull() string { return timeZoneFull(globalFaker.Rand) }

// TimeZoneFull will select a random full timezone string
func (f *Faker) TimeZoneFull() string { return timeZoneFull(f.Rand) }

func timeZoneFull(r *rand.Rand) string { return getRandValue(r, []string{"timezone", "full"}) }

// TimeZoneRegion will select a random region style timezone string, e.g. "America/Chicago"
func TimeZoneRegion() string { return timeZoneRegion(globalFaker.Rand) }

// TimeZoneRegion will select a random region style timezone string, e.g. "America/Chicago"
func (f *Faker) TimeZoneRegion() string { return timeZoneRegion(f.Rand) }

func timeZoneRegion(r *rand.Rand) string { return getRandValue(r, []string{"timezone", "region"}) }

// TimeZoneAbv will select a random timezone abbreviation string
func TimeZoneAbv() string { return timeZoneAbv(globalFaker.Rand) }

// TimeZoneAbv will select a random timezone abbreviation string
func (f *Faker) TimeZoneAbv() string { return timeZoneAbv(f.Rand) }

func timeZoneAbv(r *rand.Rand) string { return getRandValue(r, []string{"timezone", "abr"}) }

// TimeZoneOffset will select a random timezone offset
func TimeZoneOffset() float32 { return timeZoneOffset(globalFaker.Rand) }

// TimeZoneOffset will select a random timezone offset
func (f *Faker) TimeZoneOffset() float32 { return timeZoneOffset(f.Rand) }

func timeZoneOffset(r *rand.Rand) float32 {
	value, _ := strconv.ParseFloat(getRandValue(r, []string{"timezone", "offset"}), 32)
	return float32(value)
}

func addDateTimeLookup() {
	AddFuncLookup("date", Info{
		Display:     "Date",
		Category:    "time",
		Description: "Random date",
		Example:     "2006-01-02T15:04:05Z07:00",
		Output:      "string",
		Params: []Param{
			{
				Field:       "format",
				Display:     "Format",
				Type:        "string",
				Default:     "RFC3339",
				Options:     []string{"ANSIC", "UnixDate", "RubyDate", "RFC822", "RFC822Z", "RFC850", "RFC1123", "RFC1123Z", "RFC3339", "RFC3339Nano"},
				Description: "Date time string format output",
			},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			format, err := info.GetString(m, "format")
			if err != nil {
				return nil, err
			}

			switch format {
			case "ANSIC":
				return Date().Format(time.ANSIC), nil
			case "UnixDate":
				return Date().Format(time.UnixDate), nil
			case "RubyDate":
				return Date().Format(time.RubyDate), nil
			case "RFC822":
				return Date().Format(time.RFC822), nil
			case "RFC822Z":
				return Date().Format(time.RFC822Z), nil
			case "RFC850":
				return Date().Format(time.RFC850), nil
			case "RFC1123":
				return Date().Format(time.RFC1123), nil
			case "RFC1123Z":
				return Date().Format(time.RFC1123Z), nil
			case "RFC3339":
				return Date().Format(time.RFC3339), nil
			case "RFC3339Nano":
				return Date().Format(time.RFC3339Nano), nil
			}

			return "", errors.New("Invalid format")
		},
	})

	AddFuncLookup("nanosecond", Info{
		Display:     "Nanosecond",
		Category:    "time",
		Description: "Random nanosecond",
		Example:     "196446360",
		Output:      "int",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return nanoSecond(r), nil
		},
	})

	AddFuncLookup("second", Info{
		Display:     "Second",
		Category:    "time",
		Description: "Random second",
		Example:     "43",
		Output:      "int",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return second(r), nil
		},
	})

	AddFuncLookup("minute", Info{
		Display:     "Minute",
		Category:    "time",
		Description: "Random minute",
		Example:     "34",
		Output:      "int",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return minute(r), nil
		},
	})

	AddFuncLookup("hour", Info{
		Display:     "Hour",
		Category:    "time",
		Description: "Random hour",
		Example:     "8",
		Output:      "int",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return hour(r), nil
		},
	})

	AddFuncLookup("day", Info{
		Display:     "Day",
		Category:    "time",
		Description: "Random day",
		Example:     "12",
		Output:      "int",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return day(r), nil
		},
	})

	AddFuncLookup("weekday", Info{
		Display:     "Weekday",
		Category:    "time",
		Description: "Random week day",
		Example:     "Friday",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return weekDay(r), nil
		},
	})

	AddFuncLookup("year", Info{
		Display:     "Year",
		Category:    "time",
		Description: "Random year",
		Example:     "1900",
		Output:      "int",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return year(r), nil
		},
	})

	AddFuncLookup("timezone", Info{
		Display:     "Timezone",
		Category:    "time",
		Description: "Random timezone",
		Example:     "Kaliningrad Standard Time",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return timeZone(r), nil
		},
	})

	AddFuncLookup("timezoneabv", Info{
		Display:     "Timezone Abbreviation",
		Category:    "time",
		Description: "Random abbreviated timezone",
		Example:     "KST",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return timeZoneAbv(r), nil
		},
	})

	AddFuncLookup("timezonefull", Info{
		Display:     "Timezone Full",
		Category:    "time",
		Description: "Random full timezone",
		Example:     "(UTC+03:00) Kaliningrad, Minsk",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return timeZoneFull(r), nil
		},
	})

	AddFuncLookup("timezoneoffset", Info{
		Display:     "Timezone Offset",
		Category:    "time",
		Description: "Random timezone offset",
		Example:     "3",
		Output:      "float32",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return timeZoneOffset(r), nil
		},
	})

	AddFuncLookup("timezoneregion", Info{
		Display:     "Timezone Region",
		Category:    "time",
		Description: "Random region timezone",
		Example:     "America/Alaska",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return timeZoneRegion(r), nil
		},
	})

}
