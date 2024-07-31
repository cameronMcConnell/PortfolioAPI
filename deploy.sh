# Install docker
sudo snap install docker

# Build the container
docker build -t portfolio-proxy .

# Run the container
docker run -d -p 443:8080 portfolio-proxy