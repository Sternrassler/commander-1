# ADR-0001: Linux Platform Support

## Status

Accepted

## Context

The project is being developed as a cross-platform application.
An important decision concerns the support for different operating systems and architectures.

## Decision

The project will support the following platforms in the future:

- **Linux** (x86_64, ARM64)
- **macOS** (ARM64, Apple Silicon)

### Unsupported Platforms

- **Windows**: The project does not currently support Windows and is not
  planned for future support.

## Rationale

1. **Linux Focus**: As an open-source project, Linux is the primary
   target platform for development and deployment.
2. **ARM Support**: With the growth of ARM servers and embedded devices,
   broad architecture support is important.
3. **x86 Support**: The majority of desktop and server environments
   still run on x86, so this architecture is essential.
4. **No Windows Support**: Development resources are focused on
   Linux and macOS compatibility to ensure stability and quality.

## Consequences

### Positive

- Clear focus on Unix-like environments (Linux and macOS)
- Simpler testing and deployment
- Better resource utilization for development

### Negative

- No ability for Windows users to directly use the project
- Potential exclusion of part of the potential user base

## References

- Project philosophy: Open Source and Linux-friendly
- Resource constraints require prioritization
