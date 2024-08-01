# Install docker
sudo snap install docker

# Build the container
docker build -t portfolio-api .

# Get site from repository
cp ../Portfolio/site .

# Run the container
docker run -d -p 80:8080 portfolio-api