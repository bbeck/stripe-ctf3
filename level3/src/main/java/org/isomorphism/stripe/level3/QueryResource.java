package org.isomorphism.stripe.level3;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.common.base.Functions;
import com.google.common.collect.Collections2;
import org.isomorphism.stripe.level3.index.Index;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.MediaType;
import java.util.Collection;

import static com.google.common.base.Preconditions.checkNotNull;

@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class QueryResource
{
  @GET
  public QueryResponse search(@QueryParam("q") String query)
  {
    Collection<Index.Match> matches = Index.INDEX.search(query);
    return new QueryResponse(true, Collections2.transform(matches, Functions.toStringFunction()));
  }

  public static final class QueryResponse
  {
    @JsonProperty
    public final String success;

    @JsonProperty
    public final Collection<String> results;

    private QueryResponse(boolean success, Collection<String> results)
    {
      this.success = Boolean.toString(success);
      this.results = checkNotNull(results);
    }
  }
}
