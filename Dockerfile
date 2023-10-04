FROM zhield/shell:stable

EXPOSE 9090/tcp

VOLUME ["/etc/didctl"]

COPY didctl /usr/bin/didctl
ENTRYPOINT ["/usr/bin/didctl"]
