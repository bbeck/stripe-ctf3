package org.isomorphism.stripe.level3.index;

import com.google.common.base.Objects;
import com.google.common.collect.Lists;
import com.google.common.io.CharStreams;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Path;
import java.util.Collection;
import java.util.List;

public class Index
{
  public static final Index INDEX = new Index();

  private final List<Page> pages = Lists.newArrayList();

  public void addPath(Path root, Path full) throws IOException
  {
    File file = full.toFile();
    File relative = root.relativize(full).toFile();

    List<String> lines = CharStreams.readLines(new BufferedReader(new FileReader(file)));
    pages.add(new Page(relative, lines));
  }

  public Collection<Match> search(String query)
  {
    Collection<Match> matches = Lists.newArrayList();
    for (Page page : pages) {
      page.findAll(query, matches);
    }

    return matches;
  }

  /** A page is an encapsulation of ~4k of bytes from a file. */
  private static final class Page
  {
    /** Which file the data is from */
    private final File source;

    /** The lines from the file. */
    private final Collection<String> lines;

    Page(File source, Collection<String> lines)
    {
      this.source = source;
      this.lines = lines;
    }

    void findAll(String query, Collection<Match> matches)
    {
      int lineno = 0;
      for (String line : lines) {
        if (line.contains(query)) {
          matches.add(new Match(source, lineno));
        }
        lineno++;
      }
    }

    @Override
    public String toString()
    {
      return Objects.toStringHelper(this)
          .add("source", source)
          .toString();
    }
  }

  public static final class Match
  {
    public final File source;
    public final int offset;

    Match(File source, int offset)
    {
      this.source = source;
      this.offset = offset + 1;
    }

    @Override
    public String toString()
    {
      return source.toString() + ":" + offset;
    }
  }
}
