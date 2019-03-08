# Padrões de estrutura

    Nome da empresa | Linguagem | Biblioteca ou Projeto

    Variaveis de ambiente

        go env
        
            GOPATH=[...]/go
            GOBIN=[...]/go/bin

        arvore $GOPATH

            [...]/go
                `-- bin
                    `-- guru
                     -- godef
                     -- godoc
                     -- etc
                `-- pkg
                    `-- linux_amd64
                        `-- github.com
                         -- gorilla
                         -- Shopify
                `-- src
                    `-- Inlog.Go.NOME_DO_PROJETO_1
                     -- Inlog.Go.NOME_DO_PROJETO_2

    Inlog.Go.NOME_DO_PROJETO

        [-pkg >> pacotes organizados por suas origens
        `-- pasta
            `-- codigo X

        [-src >> subdiretórios similares para o repositório de origin ou site
        `-- Inlog.Go.Projeto.XXX
            `-- entity/interface
             -- service
             -- command        

        [-misc >> miscellaneous (diversos)      
        `-- help
         -- samples
         -- docs
            `-- documento.xxx
         -- scripts
            `-- scripts.sql    
         -- lib
            `-- lib_bla           

        [-bin >> comandos/ferramentas do Go
        `-- golint.exe 
         -- godev.exe 

        [-test
        `-- pasta
            `-- codigo X
             -- codigo y

        [ README.md >> informações úteis para rodar a aplicação, como: variaveis de ambiente, caminohs, pastas, configurações, etc

# Versionamento

    https://semver.org/lang/pt-BR/

# Comandos

    https://golang.org/cmd/go/
    
# Testes 

    BDD
        GoConvey    
        https://blog.codeship.com/implementing-a-bdd-workflow-in-go/
        https://github.com/smartystreets/goconvey
        https://github.com/smartystreets/goconvey/wiki

## Players de mercado

    Links:

## Benchmark C# vs GO

    Performance
    Memoria 

# Message brokers

    ## Verificar:
        provider c#
        provider goLang
        envio de mensagens do Delphi       
    
    ## ActiveMQ

        Stomp: "github.com/go-stomp/stomp"

    ## RabbitMQ

        Libs: github.com/streadway/amqp

    ## Kafka

        Libs: ...

    ## NATs

        Libs: github.com/nats-io/go-nats

    ## Crossbar

        Libs: ...
        ..

# Protocolos

    ## Protocolo de comunicação

        ProtoBuf

    ## Protocolo de envio de dados

        gRPC (Alternativa a API Json)


# Bancos de dados


    ## Firebird

        Libs: 
            "github.com/nakagami/firebirdsql"

    ## Mysql

        Libs: 
            "github.com/go-sql-driver/mysql"

    ## Redis

        Libs: ...

    ## Mongo

        Libs: ...

# Logging


# Configuração


# Gerenciador de dependencias 

    https://github.com/golang/dep

# Sites 

    https://medium.freecodecamp.org/here-are-some-amazing-advantages-of-go-that-you-dont-hear-much-about-1af99de3b23a
    https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/
    

# Repositorios importantes 

    https://github.com/golang/go
    https://github.com/kubernetes/kubernetes
    https://github.com/nsqio/nsq
    https://github.com/tmrts/go-patterns
    https://github.com/go-br
    https://github.com/go-br/estudos
    https://github.com/go-br/crud