# The .NET 10 SDK with the NuGet packages of Compat.csproj restored for each provider: the EF Core
# harness builds against them offline at run time, with the generated code in Models/.
FROM mcr.microsoft.com/dotnet/sdk:10.0
ENV DOTNET_CLI_TELEMETRY_OPTOUT=1 DOTNET_NOLOGO=1
WORKDIR /harness
COPY Compat.csproj ./
RUN for p in postgresql mysql sqlite; do dotnet restore -p:HekirekiProvider=$p; done
