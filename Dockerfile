FROM appstream-cr-cn-hangzhou-registry.cn-hangzhou.cr.aliyuncs.com/appstream/appstream_runtime:cicd.docker_master.98.20251117004909

ENV PATH /usr/local/bin:$PATH

# runtime dependencies
RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		gnupg \
		netbase \
		wget \
        git \
		mercurial \
		openssh-client \
		subversion \
        procps \
        autoconf \
		automake \
		bzip2 \
		default-libmysqlclient-dev \
		dpkg-dev \
		file \
		g++ \
		gcc \
		imagemagick \
		libbz2-dev \
		libc6-dev \
		libcurl4-openssl-dev \
		libdb-dev \
		libevent-dev \
		libffi-dev \
		libgdbm-dev \
		libglib2.0-dev \
		libgmp-dev \
		libjpeg-dev \
		libkrb5-dev \
		liblzma-dev \
		libmagickcore-dev \
		libmagickwand-dev \
		libmaxminddb-dev \
		libncurses5-dev \
		libncursesw5-dev \
		libpng-dev \
		libpq-dev \
		libreadline-dev \
		libsqlite3-dev \
		libssl-dev \
		libtool \
		libwebp-dev \
		libxml2-dev \
		libxslt-dev \
		libyaml-dev \
		make \
		patch \
		unzip \
		xz-utils \
		zlib1g-dev \
	; \
	rm -rf /var/lib/apt/lists/*

ENV NODE_VERSION=25.2.1