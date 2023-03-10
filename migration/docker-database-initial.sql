create table animals(
    id         		serial primary key,
	sex				varchar	,
    datbirth 		varchar, 
	birthweight		integer	,
	flock  	        varchar	,
	racename		varchar ,
	weightnow		integer ,
	datesale 		varchar,
	saleweight		integer ,
	situation		varchar	
);

create table actions(
    id                  serial primary key,
    Actiondate 		    varchar,
	flock		        varchar ,
	Descriptions  	    varchar	,
	Affectedanimals		varchar ,
	Cost		        integer ,
	Vermicide 		    integer ,
	Supplementation		integer 
	
);

create table reports(
    id	 	 	                serial primary key,
	Problemcode					integer ,
    Datephysicalproblem		    varchar,
	Descriptions 		        varchar	,
	Imageproblem	 	  	    varchar	

);