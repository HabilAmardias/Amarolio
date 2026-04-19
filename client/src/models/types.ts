export interface Profile {
  name: string;
  title: string;
  bio: string;
  avatarUrl: string;
  socials: { label: string; url: string; icon: string }[];
}

export interface Experience {
  id: string;
  company: string;
  role: string;
  period: string;
  location: string;
  description: string;
  tags: string[];
}

export interface Project {
  id: string;
  title: string;
  description: string;
  tags: string[];
  imageUrl: string;
  liveUrl?: string;
  repoUrl?: string;
  featured: boolean;
}
