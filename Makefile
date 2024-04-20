sizes = 10 100 1000 10000 100000 1000000

.PHONY: all
all: $(sizes) bin/report
	@ echo "generating summary report..."
	@ bin/report $(foreach var,$(sizes),-file bin/timing-$(var).txt)

bin/create_measurements: $(wildcard cmd/create_measurements/*.go)
	@ go build -o bin/create_measurements ./cmd/create_measurements

bin/calculate_average: $(wildcard cmd/calculate_average/*.go)
	@ go build -o bin/calculate_average ./cmd/calculate_average

bin/report: $(wildcard cmd/report/*.go)
	@ go build -o bin/report ./cmd/report

$(sizes): bin/create_measurements bin/calculate_average
	@ echo "creating $@ measurements..."
	@ bin/create_measurements -file bin/measurements-$@.txt -size $@
	@ echo "calculating averages for $@ measurements..."
	@ bin/calculate_average -in bin/measurements-$@.txt -result bin/result-$@.txt -timing bin/timing-$@.txt

.PHONY: clean
clean:
	@ rm -rf bin
