# Get site from repository
cp -r ../Portfolio/site .

# Build the container
docker build -t portfolio-api .

# Run the container
docker run -d -p 80:80 portfolio-api