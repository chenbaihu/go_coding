# Configs
MUDUO_INC_PATH = ./ 
MUDUO_LIB_PATH = ./muduo/lib

PWD := $(shell pwd)

# Templates
CXX := g++
DEBUG=-D_DEBUG -O0 
CXXFLAGS = -g3 $(DEBUG) -fPIC -DMUDUO_STD_STRING \
		   -Wall -Werror -MMD\
		   -Wno-unused-parameter -Wno-unused-function -Wunused-variable -Wunused-value \
		   -Wcast-qual -Wcast-align -Wwrite-strings -Wsign-compare \
		   -Winvalid-pch -fms-extensions -Wfloat-equal -Wextra \
		   -I./ \
		   -I$(MUDUO_INC_PATH) \
		   -I/usr/local/include  

LDFLAGS := -lpthread \
	-L/usr/local/lib \

LIBS := $(MUDUO_LIB_PATH)/libmuduo_net_udp.a \
	$(MUDUO_LIB_PATH)/libmuduo_net.a \
	$(MUDUO_LIB_PATH)/libmuduo_base.a \
	-lpthread

DIR := . 
SRC := $(foreach d, $(DIR), $(wildcard $(d)/*.cc))
OBJ := $(patsubst %.cc, .obj/%.o, $(SRC))
DEP := $(patsubst %.o, %.d, $(OBJ))

all: prepare tcp_echo_server 

tcp_echo_server: $(OBJ) 
	$(CXX) $^ -o $@ $(RTFLAGS) $(LDFLAGS) $(LIBS) 

prepare : 
	@mkdir -p .obj/

-include $(DEP)

.obj/%.o : %.cc
	$(CXX) -c $(CXXFLAGS) $< -o $@

clean:
	-rm -rf .obj tcp_echo_server 

test: shared
	$(MAKE) test -C ut

.PHONY: all install clean test shared

