export namespace models {
	
	export class Faction {
	    id: number;
	    name: string;
	    description?: string;
	
	    static createFrom(source: any = {}) {
	        return new Faction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class Affiliation {
	    id: number;
	    name: string;
	    description?: string;
	    type?: string;
	    faction_id?: number;
	    faction?: Faction;
	
	    static createFrom(source: any = {}) {
	        return new Affiliation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.type = source["type"];
	        this.faction_id = source["faction_id"];
	        this.faction = this.convertValues(source["faction"], Faction);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Alignment {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Alignment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Article {
	    id: number;
	    title: string;
	    description: string;
	    category: string;
	    tags?: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Article(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.category = source["category"];
	        this.tags = source["tags"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class LoreCategory {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new LoreCategory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Lore {
	    id: number;
	    title: string;
	    content: string;
	    category_id: number;
	    character_id?: number;
	    // Go type: time
	    created_at: any;
	    lore_category?: LoreCategory;
	    character?: Character;
	
	    static createFrom(source: any = {}) {
	        return new Lore(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.content = source["content"];
	        this.category_id = source["category_id"];
	        this.character_id = source["character_id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.lore_category = this.convertValues(source["lore_category"], LoreCategory);
	        this.character = this.convertValues(source["character"], Character);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Relationship {
	    id: number;
	    source_char_id: number;
	    target_char_id: number;
	    type: string;
	    notes?: string;
	    source_character?: Character;
	    target_character?: Character;
	
	    static createFrom(source: any = {}) {
	        return new Relationship(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.source_char_id = source["source_char_id"];
	        this.target_char_id = source["target_char_id"];
	        this.type = source["type"];
	        this.notes = source["notes"];
	        this.source_character = this.convertValues(source["source_character"], Character);
	        this.target_character = this.convertValues(source["target_character"], Character);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Fear {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Fear(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Weakness {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Weakness(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Flaw {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Flaw(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Strength {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Strength(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class PersonalityTrait {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new PersonalityTrait(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Location {
	    id: number;
	    name: string;
	    type: string;
	    subtype?: string;
	    description?: string;
	    parent_id?: number;
	    parent?: Location;
	    children?: Location[];
	
	    static createFrom(source: any = {}) {
	        return new Location(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.subtype = source["subtype"];
	        this.description = source["description"];
	        this.parent_id = source["parent_id"];
	        this.parent = this.convertValues(source["parent"], Location);
	        this.children = this.convertValues(source["children"], Location);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Status {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Status(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Occupation {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Occupation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Race {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Race(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Gender {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new Gender(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class Character {
	    id: number;
	    firstname: string;
	    middlename?: string;
	    surname: string;
	    nickname?: string;
	    age: string;
	    description?: string;
	    portrait?: string;
	    goals?: string;
	    gender_id?: number;
	    race_id?: number;
	    occupation_id?: number;
	    alignment_id?: number;
	    status_id?: number;
	    affiliation_id?: number;
	    faction_id?: number;
	    birthplace_location_id?: number;
	    current_location_id?: number;
	    gender?: Gender;
	    race?: Race;
	    occupation?: Occupation;
	    alignment?: Alignment;
	    status?: Status;
	    affiliation?: Affiliation;
	    faction?: Faction;
	    birthplace?: Location;
	    current_location?: Location;
	    personality_traits?: PersonalityTrait[];
	    strengths?: Strength[];
	    flaws?: Flaw[];
	    weaknesses?: Weakness[];
	    fears?: Fear[];
	    relationships?: Relationship[];
	    backstories?: Lore[];
	
	    static createFrom(source: any = {}) {
	        return new Character(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.firstname = source["firstname"];
	        this.middlename = source["middlename"];
	        this.surname = source["surname"];
	        this.nickname = source["nickname"];
	        this.age = source["age"];
	        this.description = source["description"];
	        this.portrait = source["portrait"];
	        this.goals = source["goals"];
	        this.gender_id = source["gender_id"];
	        this.race_id = source["race_id"];
	        this.occupation_id = source["occupation_id"];
	        this.alignment_id = source["alignment_id"];
	        this.status_id = source["status_id"];
	        this.affiliation_id = source["affiliation_id"];
	        this.faction_id = source["faction_id"];
	        this.birthplace_location_id = source["birthplace_location_id"];
	        this.current_location_id = source["current_location_id"];
	        this.gender = this.convertValues(source["gender"], Gender);
	        this.race = this.convertValues(source["race"], Race);
	        this.occupation = this.convertValues(source["occupation"], Occupation);
	        this.alignment = this.convertValues(source["alignment"], Alignment);
	        this.status = this.convertValues(source["status"], Status);
	        this.affiliation = this.convertValues(source["affiliation"], Affiliation);
	        this.faction = this.convertValues(source["faction"], Faction);
	        this.birthplace = this.convertValues(source["birthplace"], Location);
	        this.current_location = this.convertValues(source["current_location"], Location);
	        this.personality_traits = this.convertValues(source["personality_traits"], PersonalityTrait);
	        this.strengths = this.convertValues(source["strengths"], Strength);
	        this.flaws = this.convertValues(source["flaws"], Flaw);
	        this.weaknesses = this.convertValues(source["weaknesses"], Weakness);
	        this.fears = this.convertValues(source["fears"], Fear);
	        this.relationships = this.convertValues(source["relationships"], Relationship);
	        this.backstories = this.convertValues(source["backstories"], Lore);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	
	
	
	
	
	
	
	
	

}

export namespace services {
	
	export class LookupItem {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new LookupItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}
	export class TraitItem {
	    id: number;
	    label: string;
	
	    static createFrom(source: any = {}) {
	        return new TraitItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	    }
	}

}

