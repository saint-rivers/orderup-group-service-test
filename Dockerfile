FROM postgres:14.4-alpine
ENV PGDATA=/var/lib/postgresql/data

HEALTHCHECK --interval=10s --timeout=5s --start-period=5s --retries=3 \
    CMD ["pg_isready", "-U", "postgres"]