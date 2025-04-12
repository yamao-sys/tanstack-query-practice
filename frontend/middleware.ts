import { NextResponse, type NextRequest } from "next/server";
import { setCsrfToken } from "./apis/clients/base";

export async function middleware(request: NextRequest) {
  // NOTE: CSRFトークンをクッキーにセット
  await setCsrfToken();

  return NextResponse.next({
    request: {
      headers: request.headers,
    },
  });
}

// NOTE: ビルドファイル, 画像最適化エンドポイント, ファビコン, 画像は対象外
export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)"],
};
