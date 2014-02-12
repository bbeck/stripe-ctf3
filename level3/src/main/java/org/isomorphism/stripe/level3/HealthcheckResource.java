package org.isomorphism.stripe.level3;

import com.fasterxml.jackson.annotation.JsonProperty;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("/healthcheck")
@Produces(MediaType.APPLICATION_JSON)
public class HealthcheckResource
{
  @GET
  public HealthCheckResponse isHealthy()
  {
    return HealthCheckResponse.READY;
  }

  public static final class HealthCheckResponse
  {
    private static final HealthCheckResponse READY = new HealthCheckResponse(true);
    private static final HealthCheckResponse NOT_READY = new HealthCheckResponse(false);

    @JsonProperty("success")
    public final String ready;

    private HealthCheckResponse(boolean ready)
    {
      this.ready = Boolean.toString(ready);
    }
  }
}
