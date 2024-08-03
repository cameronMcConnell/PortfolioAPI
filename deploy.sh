# Build the container
docker build -t portfolio-api .

# Get site from repository
cp -r ../Portfolio/site .

# Run the container
docker run -d -p 80:80 portfolio-api