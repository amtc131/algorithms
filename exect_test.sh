#!/bin/bash

export LOG_TEST=$(basename $0 .sh).log
export LIST_FILE=list.tmp

if [ -f ${LOG_TEST} ] ; then
    rm ${LOG_TEST} 2>/dev/null
fi

if [ -z ${DIR_PROG} ] ; then
	if [ "`echo $0|awk -F\/ '{print NF}'`" -gt 1 ] ; then
	    export DIR_PROG="`echo $0|sed \"s/\/\`echo $0$|awk -F\/ '{print $NF}'\`//g\"`"
	fi
	if [ "${DIR_PROG}" = "./" -o "${DIR_PROG}" = "././" -o "${DIR_PROG}" = "."  ] ; then
	    export DIR_PROG=`pwd`
	fi
	if [ -z "${DIR_PROG}" ] ; then
	    export DIR_PROG=`pwd`
	fi
cd "${DIR_PROG}"
fi

exec(){
    echo "Search files test *.go" >> $LOG_TEST

    ls -c1 $DIR_PROG/test/*.go > ${LIST_FILE} 2>/dev/null

    while read tests;
    do
        NAME_TEST=${tests##*/}
        go test $DIR_PROG/test/$NAME_TEST -v >>$LOG_TEST

    done < ${LIST_FILE}
    #FAIL
    err=`grep "FAIL: Test\/" $LOG_TEST | wc -l`
    ok=`awk '/ok  / {print $0}' exect_test.log  | wc -l`
    echo "`grep "FAIL: Test\/" exect_test.log`"

    echo "Test execute"
    echo "------------"
    echo "ERROR: $err"
    echo "OK: $ok"


    echo "End Test" >> $LOG_TEST

}

main(){
    exec
}

echo "Init RUN Test" >> $LOG_TEST
main
