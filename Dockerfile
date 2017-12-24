FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY go-learning-kubernetes /
CMD ["/go-learning-kubernetes", "-port", "8000"]
