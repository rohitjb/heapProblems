FROM instrumenta/conftest as conftestInstaller
COPY /debian9 debian9
COPY /policy policy
RUN conftest test debian9/Dockerfile -o json > result.json
RUN echo result.json


