package org.isomorphism.stripe.level3;

import com.yammer.dropwizard.Service;
import com.yammer.dropwizard.config.Bootstrap;
import com.yammer.dropwizard.config.Environment;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import java.io.IOException;

public class Level3Service extends Service<Level3Configuration>
{
  public static void main(String[] args) throws Exception
  {
    new Level3Service().run(args);
  }

  @Override
  public void initialize(Bootstrap<Level3Configuration> bootstrap)
  {
  }

  @Override
  public void run(Level3Configuration configuration, Environment env) throws Exception
  {
    env.addResource(HealthcheckResource.class);
    env.addResource(IndexResource.class);
    env.addResource(IsIndexedResource.class);
    env.addResource(QueryResource.class);
  }
}
