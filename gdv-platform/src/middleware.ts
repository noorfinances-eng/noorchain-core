import { NextRequest, NextResponse } from "next/server";

type Role = "curator" | "dev";

function getRole(req: NextRequest): Role | null {
  const v = req.cookies.get("gdv_role")?.value ?? "";
  if (v === "curator" || v === "dev") return v;
  return null;
}

function isAllowed(pathname: string, role: Role): boolean {
  if (role === "dev") return true;
  // curator
  return pathname === "/curator" || pathname.startsWith("/curator/");
}

export function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  // Allow Next internals and static assets
  if (
    pathname.startsWith("/_next") ||
    pathname.startsWith("/favicon.ico") ||
    pathname.startsWith("/robots.txt") ||
    pathname.startsWith("/sitemap.xml")
  ) {
    return NextResponse.next();
  }

  // Always allow login page
  if (pathname === "/login") return NextResponse.next();

  const role = getRole(req);

  // Not authenticated -> redirect to /login
  if (!role) {
    const url = req.nextUrl.clone();
    url.pathname = "/login";
    url.searchParams.set("next", pathname);
    return NextResponse.redirect(url);
  }

  // Authenticated but not allowed -> redirect to their home
  if (!isAllowed(pathname, role)) {
    const url = req.nextUrl.clone();
    url.pathname = role === "dev" ? "/dev" : "/curator";
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/((?!api).*)"],
};
