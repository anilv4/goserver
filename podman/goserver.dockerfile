# Step 1: Build the application
# Use the official Fedora image as the build environment.
FROM registry.fedoraproject.org/fedora:latest as builder

# Enable fastest mirror in dnf configuration.
RUN echo 'fastestmirror=1' >> /etc/dnf/dnf.conf

# Install Git and Go.
RUN dnf -y update && dnf -y install git golang

# Set the working directory inside the container.
WORKDIR /app

# Clone the specific repository.
RUN git clone https://github.com/anilv4/goserver.git .

# Change to the source directory.
WORKDIR /app/source

# Build the Go app.
RUN go build -o goserver ./goserver.go

# Step 2: Use a lightweight base image for the runtime.
FROM registry.fedoraproject.org/fedora:latest

# Set environment variables with default values
ENV APP_PORT=8080
ENV APP_HOME=/tmp/home
ENV APP_LOGGING=true

# Set the working directory in the new container.
WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/source/goserver .

# Install troubleshooting tools. Ignore if you dont want it.
RUN echo 'fastestmirror=1' >> /etc/dnf/dnf.conf
RUN dnf install -y --nodocs telnet tcpdump curl util-linux lsof strace iproute net-tools nmap bind-utils procps-ng nc bash-completion jq iputils sysstat fio && dnf clean all

# Command to run the executable with environment variables.
CMD ["sh", "-c", "./goserver --port=$APP_PORT --home=$APP_HOME --logging=$APP_LOGGING"]
