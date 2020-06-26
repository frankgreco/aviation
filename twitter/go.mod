module github.com/frankgreco/aviation/twitter

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/aws/aws-lambda-go v1.17.0
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/frankgreco/aviation v0.0.0
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/utils v0.0.0-20200624030258-6c014fac6407
	github.com/jmoiron/sqlx v1.2.0
	github.com/kyokomi/emoji v2.2.4+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/yuin/goldmark v1.1.32 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200622214017-ed371f2e16b4 // indirect
	golang.org/x/text v0.3.3
	golang.org/x/tools v0.0.0-20200623204733-f8e0ea3a3a8f // indirect
)

replace (
	github.com/frankgreco/aviation => ../
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/utils => ../utils
)
