import { Helmet } from "react-helmet-async";
import { siteConfig } from "../../config/site";

interface SeoProps {
  title: string;
  description: string;
  path: string;
  type?: "website" | "profile" | "article";
  keywords?: string[];
  image?: string;
  jsonLd?: Record<string, unknown>[];
}

export default function Seo({
  title,
  description,
  path,
  type = "website",
  keywords,
  image,
  jsonLd = [],
}: SeoProps) {
  const fullTitle = title.includes(siteConfig.name)
    ? title
    : `${title} | ${siteConfig.name}`;
  const url = `${siteConfig.url}${path === "/" ? "" : path}`;
  const absoluteImage = (src: string) =>
    src.startsWith("http") ? src : `${siteConfig.url}${src}`;
  const ogImage = absoluteImage(image ?? siteConfig.image);

  return (
    <Helmet>
      <title>{fullTitle}</title>
      <meta name="description" content={description} />
      {keywords && keywords.length > 0 && (
        <meta name="keywords" content={keywords.join(", ")} />
      )}
      <link rel="canonical" href={url} />

      <meta property="og:type" content={type} />
      <meta property="og:site_name" content={siteConfig.name} />
      <meta property="og:title" content={fullTitle} />
      <meta property="og:description" content={description} />
      <meta property="og:url" content={url} />
      <meta property="og:image" content={ogImage} />
      <meta property="og:locale" content={siteConfig.locale} />

      <meta name="twitter:card" content="summary_large_image" />
      <meta name="twitter:title" content={fullTitle} />
      <meta name="twitter:description" content={description} />
      <meta name="twitter:image" content={ogImage} />

      {jsonLd.map((data, index) => (
        <script key={index} type="application/ld+json">
          {JSON.stringify(data)}
        </script>
      ))}
    </Helmet>
  );
}
