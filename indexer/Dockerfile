# Oficial image
FROM public.ecr.aws/zinclabs/zincsearch:latest

# Set environment variables
ENV ZINC_DATA_PATH=/data
ENV ZINC_FIRST_ADMIN_USER=admin
ENV ZINC_FIRST_ADMIN_PASSWORD=Complexplass#123

# Expose port 4080
EXPOSE 4080

# Add a volume for persistent data storage
VOLUME /data

LABEL com.docker.compose.container_name=zincsearch
