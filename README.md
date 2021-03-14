# Abghand
Just another simple TCP proxy that works on HTTP and HTTPS flows.

To use Abghand, first you need to send your traffic to Abghand by setting DNS records in `/etc/hosts`. Then define the domains that you want to proxy their traffic in the config file.

This is a sample of Abghand config file:  

    ---
    - hostname: 'index.docker.io'
    - hostname: 'registry.hub.docker.com'
    - hostname: 'auth.docker.io'
    - hostname: 'hub.docker.com'
    - hostname: 'registry-1.docker.io'
    - hostname: 'service-with-non-standard-port.com'
      port_set:
        - port_number: 8080
          type: "HTTP"
        - port_number: 11443
          type: "HTTPS"


#### Setting up:

To build and install from source code run the following commands:  

    go get 'github.com/Yaser-Amiri/abghand'
    go build
    go install

To run Abghand on Docker:  

    docker run --name abghand -d \
        -v $(pwd)/config.yml:/config.yml \
        -p 80:80 -p 443:443 yaseramiri/abghand
*Add or remove published ports based on your configuration.*
