create table animals(
    id         serial primary key,
	sex				varchar	,
    datbirth 		integer ,
	birthweight		integer	,
	flock  	        varchar	,
	racename		varchar ,
	weightnow		integer ,
	datesale 		integer ,
	saleweight		integer ,
	situation		varchar	
);

create table actions(
    id                  serial primary key,
	Identifier        	varchar,
    Actiondate 		    integer ,
	flock		        varchar ,
	Descriptions  	    varchar	,
	Affectedanimals		varchar ,
	Cost		        integer ,
	Vermicide 		    integer ,
	Supplementation		integer 
	
);

create table reports(
    id	 	 	                serial primary key,
    Datephysicalproblem		    integer ,
	Descriptions 		        varchar	,
	Imageproblem	 	  	    varchar	
);