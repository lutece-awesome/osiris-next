MAIN := main
SRC := .
DIST := dist
CXX := clang++-12
CXXFLAGS := -std=c++0x -o2 -Wall
SRC_CC := $(shell find ${SRC} -name '*.cc' | cut -b 3-)
OBJ_CC := $(patsubst %.cc,%.o, $(addprefix ${DIST}/, ${SRC_CC} ))

.PHONY: clean

${MAIN}: $(OBJ_CC)
	${CXX} $(CXXFLAGS) -o ${DIST}/${MAIN} $^

${DIST}/%.o: ${SRC}/%.cc
	@ mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) $< -c -o $@

clean:
	@ rm -r ${DIST}

format:
	@ clang-format -style=Google -i ${SRC_CC}