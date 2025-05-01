"use client"

import type React from "react"
import { useState, useRef } from "react"
import Image from "next/image"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { useRouter } from "next/navigation"
import { ArrowLeft } from "lucide-react"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"


interface ProfileData {
  fullName: string
  title: string
  photoUrl: string
  phoneNumber: string
  company: {
    name: string
    industry: string
    location: string
    logoUrl: string
  }
}

export default function EditProfilePerekrut() {
  const router = useRouter()
  const fileInputRef = useRef<HTMLInputElement>(null)
  const logoInputRef = useRef<HTMLInputElement>(null)

  const [profile, setProfile] = useState<ProfileData>({
    fullName: "Aditya Pratama",
    title: "Senior HR in Shopee Indonesia",
    photoUrl: "/placeholder.svg?height=120&width=120",
    phoneNumber: "081548797755",
    company: {
      name: "Shopee",
      industry: "E-Commerce",
      location: "Jakarta Pusat",
      logoUrl: "/placeholder.svg?height=80&width=80",
    },
  })

  const [isSaving, setIsSaving] = useState(false)
  const [photoPreview, setPhotoPreview] = useState<string | null>(null)
  const [logoPreview, setLogoPreview] = useState<string | null>(null)

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
  
    setProfile((prev) => {
      if (name.includes(".")) {
        const [parent, child] = name.split(".")
        const parentObj = (prev[parent as keyof ProfileData] ?? {}) as Record<string, any>
        return {...prev, [parent]: {
            ...parentObj,
            [child]: value,
          },
        }
      } else {
        return {...prev, [name]: value,
        }
      }
    })
  }

  const handlePhotoClick = () => {
    fileInputRef.current?.click()
  }

  const handleLogoClick = () => {
    logoInputRef.current?.click()
  }

  const handlePhotoChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0]
      const reader = new FileReader()

      reader.onload = (event) => {
        if (event.target?.result) {
          setPhotoPreview(event.target.result as string)
        }
      }

      reader.readAsDataURL(file)
    }
  }

  const handleLogoChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0]
      const reader = new FileReader()

      reader.onload = (event) => {
        if (event.target?.result) {
          setLogoPreview(event.target.result as string)
        }
      }

      reader.readAsDataURL(file)
    }
  }

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="perekrut"/>

      <div className="container mx-auto px-4 py-6">
        {/* Back button */}
        <div className="mb-4">
          <button onClick={() => router.back()} className="flex items-center text-[#1e3a5f] hover:underline">
            <ArrowLeft className="h-4 w-4 mr-1" />
            Kembali
          </button>
        </div>

        <h1 className="text-2xl font-bold text-center text-[#1e3a5f] mb-6">Ubah Profil</h1>

        <form className="max-w-2xl mx-auto mb-10">
          {/* Recruiter Information */}
          <div className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] p-6 mb-6">
            <h2 className="text-lg font-bold text-center text-[#1e3a5f] mb-4">Informasi Rekruter</h2>

            <div className="space-y-4">
              <div className="flex flex-col items-center mb-4">
                <div className="relative h-24 w-24 rounded-full overflow-hidden border border-gray-200 mb-2">
                  <Image src={photoPreview || profile.photoUrl} alt={profile.fullName} fill className="object-cover" />
                </div>
                <button
                  type="button"
                  onClick={handlePhotoClick}
                  className="text-sm text-[#1e3a5f] border border-[#1e3a5f] px-3 py-1 rounded-md hover:bg-[#1e3a5f]/10"
                >
                  ubah foto profil
                </button>
                <input
                  ref={fileInputRef}
                  type="file"
                  accept="image/*"
                  onChange={handlePhotoChange}
                  className="hidden"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Nama Lengkap</label>
                <input
                  type="text"
                  name="fullName"
                  value={profile.fullName}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Jabatan</label>
                <input
                  type="text"
                  name="title"
                  value={profile.title}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Nomor Telepon</label>
                <input
                  type="tel"
                  name="phoneNumber"
                  value={profile.phoneNumber}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>
            </div>
          </div>

          {/* Company Information */}
          <div className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] p-6 mb-6">
            <h2 className="text-lg font-bold text-center text-[#1e3a5f] mb-4">Informasi Perusahaan</h2>

            <div className="space-y-4">
              <div className="flex flex-col items-center mb-4">
                <div className="relative h-20 w-20 bg-red-100 rounded-md p-2 flex items-center justify-center mb-2">
                  <Image
                    src={logoPreview || profile.company.logoUrl}
                    alt={profile.company.name}
                    width={60}
                    height={60}
                    className="object-contain"
                  />
                </div>
                <button
                  type="button"
                  onClick={handleLogoClick}
                  className="text-sm text-[#1e3a5f] border border-[#1e3a5f] px-3 py-1 rounded-md hover:bg-[#1e3a5f]/10"
                >
                  ubah logo perusahaan
                </button>
                <input ref={logoInputRef} type="file" accept="image/*" onChange={handleLogoChange} className="hidden" />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Nama Perusahaan</label>
                <input
                  type="text"
                  name="company.name"
                  value={profile.company.name}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Industri</label>
                <input
                  type="text"
                  name="company.industry"
                  value={profile.company.industry}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Lokasi</label>
                <input
                  type="text"
                  name="company.location"
                  value={profile.company.location}
                  onChange={handleChange}
                  required
                  className="w-full px-3 py-2 bg-white border border-indigo-900 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>
            </div>
          </div>

          <div className="flex justify-center">
            <Button
              type="submit"
              disabled={isSaving}
              className="bg-[#1e3a5f] text-white px-6 py-2 rounded-md hover:bg-[#1e3a5f]/90 transition-colors disabled:opacity-50 cursor-pointer"
            >
              {isSaving ? "Menyimpan..." : "Simpan Profil"}
            </Button>
          </div>
        </form>
      </div>

      {/* Footer */}
      <Footer />
    </div>
  )
}
