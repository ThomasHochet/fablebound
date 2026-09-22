import { FetchAllLookup, FetchAllTraits, FetchAllLocations, FetchAllAffiliations, FetchAllFactions } from "../../../bindings/world-builder/app"

export async function loadCharacterFormData() {
  const [
    genders,
    races,
    occupations,
    alignments,
    statuses,
    personalityTraits,
    strengths,
    flaws,
    weaknesses,
    fears,
    locations,
    affiliations,
    factions
  ] = await Promise.all([
    FetchAllLookup('gender'),
    FetchAllLookup('race'),
    FetchAllLookup('occupation'),
    FetchAllLookup('alignment'),
    FetchAllLookup('status'),
    FetchAllTraits('personality'),
    FetchAllTraits('strength'),
    FetchAllTraits('flaw'),
    FetchAllTraits('weakness'),
    FetchAllTraits('fear'),
    FetchAllLocations(),
    FetchAllAffiliations(),
    FetchAllFactions()
  ])

  return {
    genders,
    races,
    occupations,
    alignments,
    statuses,
    personalityTraits,
    strengths,
    flaws,
    weaknesses,
    fears,
    locations,
    affiliations,
    factions
  }
}
