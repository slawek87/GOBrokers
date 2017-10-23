FROM elasticsearch:5.6.3
EXPOSE 9200 9300

ENV discovery.type=single-node
VOLUME ./docker/data