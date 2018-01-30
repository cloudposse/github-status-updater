FROM scratch

COPY github-commit-status /app/

WORKDIR /app

CMD ["./github-commit-status"]
