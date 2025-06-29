# TOML Configuration Guide

## TOML Configuration Order Requirements

When using TOML configuration files with Traefik, it's important to understand a key requirement of the TOML format:

!!! warning "Important TOML Requirement"
    **Non-table values must be placed before any table declarations in TOML files.**

### Correct Configuration Order
