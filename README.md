## mars agent


#### 构建
```
make VERSION=0.0.1 all
cd bin/osx/
./mars_agent -version

#ouput:
mars agent 
git_tag:, version:0.0.1, build_time: (2019-06-21T14:19:48+0800)

./mars_agent -h

#output:
mars agent version: 0.0.2
Usage: mars_agent  [-s name] [-version] [-h]

Options:
  -h    this help
  -s name
        start consul mode by name : gor, jmeter
  -version
        print agent version information
```

#### 运行
```
./mars_agent -s gor

or

./mars_agent -s jmeter
```