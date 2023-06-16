# Use the official MongoDB base image
FROM mongodb:6.0

# Set a custom data directory
ENV MONGO_DATA_DIR /data/db

# Create the data directory
RUN mkdir -p $MONGO_DATA_DIR

# Set permissions for the data directory
RUN chown -R mongodb:mongodb $MONGO_DATA_DIR

# Set the command to run MongoDB
CMD ["mongod"]

# Expose the default MongoDB port
EXPOSE 27017

#
#docker run -d -p 27017:27017 --name mongodb-container mongodb

