package org.isomorphism.stripe.level3;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.yammer.dropwizard.config.Configuration;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;

public class Level3Configuration extends Configuration {
  @JsonProperty
  @Valid
  @NotNull
  private final String id = null;
}
