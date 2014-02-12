package org.isomorphism.stripe.level3;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.isomorphism.stripe.level3.index.Index;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.MediaType;
import java.io.IOException;
import java.nio.file.FileVisitResult;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.SimpleFileVisitor;
import java.nio.file.attribute.BasicFileAttributes;

@Path("/index")
@Produces(MediaType.APPLICATION_JSON)
public class IndexResource
{
  private static final Logger LOG = LoggerFactory.getLogger(IndexResource.class);

  @GET
  public IndexResponse index(@QueryParam("path") String path)
  {
    try {
      final java.nio.file.Path root = Paths.get(path);
      Files.walkFileTree(root, new SimpleFileVisitor<java.nio.file.Path>() {
        @Override
        public FileVisitResult visitFile(java.nio.file.Path file, BasicFileAttributes attrs) throws IOException
        {
          Index.INDEX.addPath(root, file);
          return FileVisitResult.CONTINUE;
        }
      });
    } catch (IOException e) {
      LOG.warn("Exception while walking the file tree", e);
      return IndexResponse.FAILURE;
    }

    return IndexResponse.SUCCESS;
  }

  public static final class IndexResponse
  {
    private static final IndexResponse SUCCESS = new IndexResponse(true);
    private static final IndexResponse FAILURE = new IndexResponse(false);

    @JsonProperty
    public final String success;

    private IndexResponse(boolean ready)
    {
      this.success = Boolean.toString(ready);
    }
  }
}
