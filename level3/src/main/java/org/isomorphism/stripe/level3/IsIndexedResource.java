package org.isomorphism.stripe.level3;

import com.fasterxml.jackson.annotation.JsonProperty;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("/isIndexed")
@Produces(MediaType.APPLICATION_JSON)
public class IsIndexedResource
{
  @GET
  public IsIndexedResponse isIndexed()
  {
    return IsIndexedResponse.SUCCESS;
  }

  public static final class IsIndexedResponse
  {
    private static final IsIndexedResponse SUCCESS = new IsIndexedResponse(true);
    private static final IsIndexedResponse FAILURE = new IsIndexedResponse(false);

    @JsonProperty
    public final String success;

    private IsIndexedResponse(boolean ready)
    {
      this.success = Boolean.toString(ready);
    }
  }
}