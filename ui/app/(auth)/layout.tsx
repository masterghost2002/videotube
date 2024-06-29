export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body >
        <div className="p-10">
        <h3 className="text-4xl leading-[80px] font-bold">
            Video Tube
        </h3>
        {children}
        </div>
    </body>
    </html>
  );
}
