provider "userclouds" {
  tenant_url = "https://mytenant.tenant.userclouds.tools"
  # Set USERCLOUDS_CLIENT_ID and USERCLOUDS_CLIENT_SECRET environment variables to avoid hardcoding
  # secrets. USERCLOUDS_TENANT_URL may be set in place of tenant_url as well.
}
